// =============================================================================
// show.go - 显示技能详情命令实现
// =============================================================================
// 这个文件实现了PicoClaw的技能详情显示功能。
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
// newShowCommand - 创建显示详情命令
// =============================================================================
// 创建"picoclaw skills show"子命令。
// 显示指定技能的详细信息。
//
// 参数：
//   - loaderFn: 返回技能加载器的函数
//
// 位置参数：
//   - args[0]: 技能名称
//
// 使用示例：
//
//	picoclaw skills show weather
func newShowCommand(loaderFn func() (*skills.SkillsLoader, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "show",
		Short:   "Show skill details",
		Args:    cobra.ExactArgs(1), // 必须提供1个参数（技能名称）
		Example: `picoclaw skills show weather`,
		RunE: func(_ *cobra.Command, args []string) error {
			loader, err := loaderFn()
			if err != nil {
				return err
			}
			skillsShowCmd(loader, args[0])
			return nil
		},
	}

	return cmd
}
