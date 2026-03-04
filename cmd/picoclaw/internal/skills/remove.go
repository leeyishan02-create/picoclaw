// =============================================================================
// remove.go - 移除技能命令实现
// =============================================================================
// 这个文件实现了PicoClaw的技能移除功能。
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
// newRemoveCommand - 创建移除命令
// =============================================================================
// 创建"picoclaw skills remove"子命令（别名"rm", "uninstall"）。
// 移除已安装的技能。
//
// 参数：
//   - installerFn: 返回技能安装器的函数
//
// 位置参数：
//   - args[0]: 技能名称
//
// 使用示例：
//
//	picoclaw skills remove weather
//	picoclaw skills rm weather
//	picoclaw skills uninstall weather
func newRemoveCommand(installerFn func() (*skills.SkillInstaller, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove",
		Aliases: []string{"rm", "uninstall"}, // 命令别名
		Short:   "Remove installed skill",
		Args:    cobra.ExactArgs(1), // 必须提供1个参数（技能名称）
		Example: `picoclaw skills remove weather`,
		RunE: func(_ *cobra.Command, args []string) error {
			installer, err := installerFn()
			if err != nil {
				return err
			}
			skillsRemoveCmd(installer, args[0])
			return nil
		},
	}

	return cmd
}
