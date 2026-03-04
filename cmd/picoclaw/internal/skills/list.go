// =============================================================================
// list.go - 列出技能命令实现
// =============================================================================
// 这个文件实现了PicoClaw的技能列表功能。
// =============================================================================

package skills

// =============================================================================
// 导入 (Imports)
// =============================================================================
import (
	"github.com/spf13/cobra" // CLI命令库

	"github.com/sipeed/picoclaw/pkg/skills" // 技能核心功能
)

// =============================================================================
// newListCommand - 创建列出命令
// =============================================================================
// 创建"picoclaw skills list"子命令。
// 列出所有已安装的技能（包括工作区、全局和内置技能）。
//
// 参数：
//   - loaderFn: 返回技能加载器的函数
//
// 使用示例：
//
//	picoclaw skills list
func newListCommand(loaderFn func() (*skills.SkillsLoader, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List installed skills",
		Example: `picoclaw skills list`,
		RunE: func(_ *cobra.Command, _ []string) error {
			loader, err := loaderFn()
			if err != nil {
				return err
			}
			skillsListCmd(loader)
			return nil
		},
	}

	return cmd
}
