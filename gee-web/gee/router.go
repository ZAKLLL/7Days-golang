package gee

import (
	"log"
	"net/http"
	"strings"
)

//路由模块 通过 httpMethod  - pattern  定位目标Handher

type router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*node
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc), roots: make(map[string]*node)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parsePattern(pattern), 0)
	key := method + "-" + pattern
	r.handlers[key] = handler
	log.Printf("Route %4s - %s", method, pattern)

}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if len(part) == 0 {
				continue
			}
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}

//只允许一个*
func parsePattern(pattern string) []string {
	pattern = strings.TrimSpace(pattern)
	if pattern == "/" {
		return []string{""}
	}
	vs := strings.Split(pattern, "/")
	ret := make([]string, 0)
	for index, item := range vs {
		item = strings.TrimSpace(item)
		if item != "" {
			if item[0] == '*' {
				if !(item == "*" || item == "**") || index != len(vs)-1 {
					log.Fatalf("url pattern仅支持单个 /* 和 /** 模式,并以/* 或 /** 结尾! 当前pattern: %s", pattern)
				}
			}
			ret = append(ret, item)
		}
	}
	return ret
}
