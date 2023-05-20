package geecache

import (
	"fmt"
	"geecache6/consistenthash"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const defaultBasePath = "/_geecache/"
const defaultReplicas = 50

//每一台主机（存储节点），即是一个服务端也是一个客户端
type HTTPPool struct {
	self     string // 记录自己的地址，包括主机名，IP，端口
	basePath string
	mu       sync.Mutex // guards peers and httpGetters

	peers       *consistenthash.Map    //一致性哈希算法的 Map，用来根据具体的 key 选择节点。
	httpGetters map[string]*httpGetter // keyed by e.g. "http://10.0.0.2:8008",映射远程节点与对应的httpGetter
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (p *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("{Server %s} %s", p.self, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// func HasPrefix(s, prefix string) bool :
	// HasPrefix测试字符串s是否以prefix开头
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path:" + r.URL.Path)
	}

	p.Log("%s %s", r.Method, r.URL.Path)

	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)

	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group:"+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
}

//实现PeerPicker接口
// Set updates the pool's list of peers.
func (p *HTTPPool) Set(peers ...string) {
	//实例化一致性哈希算法，并添加传入的节点
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers = consistenthash.New(defaultReplicas, nil) //创建哈希轮
	p.peers.Add(peers...)                              //实际往哈希轮中添加节点（创建相应的虚拟节点-放入哈希轮）
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers { //将
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}

// PickPeer picks a peer according to key
func (p *HTTPPool) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Pick peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}

//httpGetter 实现 客户端 的功能。实现 PeerGetter 接口。
type httpGetter struct {
	baseURL string
}

//获取返回值
func (h *httpGetter) Get(group string, key string) ([]byte, error) {
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		url.QueryEscape(group),
		url.QueryEscape(key),
	)

	//  baseURL 表示将要访问的远程节点的地址，例如 http://example.com/_geecache/。
	res, err := http.Get(u) //获取连接
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v", res.Status)
	}

	// ReadAll从r开始读取，直到出现错误或EOF，并返回所读取的数据。
	//调用成功返回err == nil，而不是err == EOF。因为ReadAll被定义为从src读取直到EOF，所以它不会将read的EOF视为要报告的错误。
	//  使用 http.Get() 方式获取返回值，并转换为 []bytes 类型。
	bytes, err := ioutil.ReadAll(res.Body) //此处传出的err不会是EOF，从res.body中读取内容
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}

	return bytes, nil
}

var _ PeerGetter = (*httpGetter)(nil)

var _ PeerPicker = (*HTTPPool)(nil)
