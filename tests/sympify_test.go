/*
 * @Author: flyyuan 740004544@qq.com
 * @Date: 2023-08-13 18:29:51
 * @LastEditors: flyyuan 740004544@qq.com
 * @LastEditTime: 2023-08-13 18:50:56
 * @FilePath: /symgo/tests/sympify_test.go
 * @Description:
 */
package symgo

import (
	"symgo"
	"testing"
)

func TestSympify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"3 + 7", "10"}, // 假设这是正确的输出
		{"invalid_expr", "Error: invalid expression"}, // 假设这是错误消息
		// 添加更多测试用例...
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := symgo.Sympify(tt.input)
			if got != tt.expected {
				t.Errorf("Sympify(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
