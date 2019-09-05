package json_test

import "testing"

/**
使用json.Marshal通用的是使用该机制解析配置文件，但是当QPS的要求很高的时候，不赞成使用这种方式，因为
json.Marshal底层使用的是反射实现，效率比较低。
使用easyjson来解析，这个包解析比较快
 */
func TestJson(t *testing.T)  {

}