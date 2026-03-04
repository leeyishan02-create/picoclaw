// =============================================================================
// search.go - 搜索技能命令实现
// =============================================================================
// 这个文件实现了PicoClaw的技能搜索功能。
// 用户可以搜索ClawHub上的技能。
// =============================================================================

package skills

// =============================================================================
// 导入 (Imports)
// =============================================================================
import (
	"github.com/spf13/cobra" // CLI命令库
)

// =============================================================================
// newSearchCommand - 创建搜索命令
// =============================================================================
// 创建"picoclaw skills search"子命令。
// 搜索ClawHub上的可用技能。
//
// 位置参数（可选）：
//   - args[0]: 搜索关键词
//
// 使用示例：
//
//	# 列出所有可用技能
//	picoclaw skills search
//
//	# 搜索特定技能
//	picoclaw skills search weather
func newSearchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search [query]",
		Short: "Search available skills",
		Args:  cobra.MaximumNArgs(1), // 最多接受1个参数
		RunE: func(_ *cobra.Command, args []string) error {
			query := ""
			if len(args) == 1 {
				query = args[0]
			}
			skillsSearchCmd(query)
			return nil
		},
	}

	return cmd
}
