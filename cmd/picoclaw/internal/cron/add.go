// =============================================================================
// add.go - 添加定时任务命令实现
// =============================================================================
// 这个文件实现了PicoClaw的添加定时任务功能。
// 用户可以创建两种类型的定时任务：
// - 间隔任务：每隔指定毫秒执行一次
// - Cron任务：按照Cron表达式执行
// =============================================================================

package cron

// =============================================================================
// 导入 (Imports)
// =============================================================================
import (
	"fmt" // 格式化输出

	"github.com/spf13/cobra" // CLI命令库

	"github.com/sipeed/picoclaw/pkg/cron" // 定时任务核心功能
)

// =============================================================================
// newAddCommand - 创建添加命令
// =============================================================================
// 创建"picoclaw cron add"子命令。
// 添加新的定时任务。
//
// 命令行参数：
//   - --name: 任务名称（必填）
//   - --message: 要发送的消息内容（必填）
//   - --every: 间隔时间（毫秒），与--cron二选一
//   - --cron: Cron表达式，与--every二选一
//   - --deliver: 是否立即发送一次
//   - --channel: 指定消息通道
//   - --to: 指定接收者
//
// 使用示例：
//
//	# 每隔1小时发送消息
//	picoclaw cron add --name "hourly-reminder" --message "喝水啦" --every 3600000
//
//	# 每天早上8点发送消息
//	picoclaw cron add --name "morning" --message "早安" --cron "0 8 * * *"
func newAddCommand(storePath func() string) *cobra.Command {
	// 定义命令行标志变量
	var (
		name    string // 任务名称
		message string // 消息内容
		every   int64  // 间隔时间（毫秒）
		cronExp string // Cron表达式
		deliver bool   // 是否立即发送
		channel string // 消息通道
		to      string // 接收者
	)

	// 创建添加命令
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new scheduled job",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// 验证参数：必须指定--every或--cron之一
			if every <= 0 && cronExp == "" {
				return fmt.Errorf("either --every or --cron must be specified")
			}

			// 创建调度配置
			var schedule cron.CronSchedule
			if every > 0 {
				// 间隔任务：每隔every毫秒执行一次
				everyMS := every * 1000
				schedule = cron.CronSchedule{Kind: "every", EveryMS: &everyMS}
			} else {
				// Cron任务：按照Cron表达式执行
				schedule = cron.CronSchedule{Kind: "cron", Expr: cronExp}
			}

			// 创建定时任务服务并添加任务
			cs := cron.NewCronService(storePath(), nil)
			job, err := cs.AddJob(name, schedule, message, deliver, channel, to)
			if err != nil {
				return fmt.Errorf("error adding job: %w", err)
			}

			// 打印成功消息
			fmt.Printf("✓ Added job '%s' (%s)\n", job.Name, job.ID)

			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Job name")
	cmd.Flags().StringVarP(&message, "message", "m", "", "Message for agent")
	cmd.Flags().Int64VarP(&every, "every", "e", 0, "Run every N seconds")
	cmd.Flags().StringVarP(&cronExp, "cron", "c", "", "Cron expression (e.g. '0 9 * * *')")
	cmd.Flags().BoolVarP(&deliver, "deliver", "d", false, "Deliver response to channel")
	cmd.Flags().StringVar(&to, "to", "", "Recipient for delivery")
	cmd.Flags().StringVar(&channel, "channel", "", "Channel for delivery")

	_ = cmd.MarkFlagRequired("name")
	_ = cmd.MarkFlagRequired("message")
	cmd.MarkFlagsMutuallyExclusive("every", "cron")

	return cmd
}
