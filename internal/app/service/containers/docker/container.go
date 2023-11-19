package docker

import (
	"bytes"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"io"
)

/*
	container.Config：
		Hostname        string              // 主机名
		Domainname      string              // 域名
		User            string              // 将在容器内运行命令的用户，还支持User:group
		AttachStdin     bool                // 附加标准输入，实现用户交互
		AttachStdout    bool                // 附加标准输出
		AttachStderr    bool                // 附加标准误差
		ExposedPorts    nat.PortSet         `json:",omitempty"` //暴露端口列表
		Tty             bool                // 将标准流附加到Tty，如果Tty未关闭，则包括stdin。
		OpenStdin       bool                // Open stdin
		StdinOnce       bool                // 如果为true，请在连接的1个客户端断开连接后关闭stdin。
		Env             []string            // 要在容器中设置的环境变量列表
		Cmd             strslice.StrSlice   // 启动容器时要运行的命令
		Healthcheck     *HealthConfig       `json:",omitempty"` // Healthcheck描述了如何检查容器是否健康
		ArgsEscaped     bool                `json:",omitempty"` // True，如果命令已转义（意味着将其视为命令行）（特定于Windows）。
		Image           string              // 操作传递的镜像的名称（例如，可以是符号）
		Volumes         map[string]struct{} // 用于容器的卷（装载）列表
		WorkingDir      string              // 将启动命令中的当前目录（PWD）
		Entrypoint      strslice.StrSlice   // 启动容器时要运行的入口点
		NetworkDisabled bool                `json:",omitempty"` // 已禁用网络
		MacAddress      string              `json:",omitempty"` // 容器的Mac地址
		OnBuild         []string            // image Dockerfile上定义的OnBuild元数据
		Labels          map[string]string   // 设置到此容器的标签列表
		StopSignal      string              `json:",omitempty"` // 停止容器的信号
		StopTimeout     *int                `json:",omitempty"` // 停止容器的超时（秒）
		Shell           strslice.StrSlice   `json:",omitempty"` // Shell表示RUN的Shell形式，CMD，ENTRYPOINT
*/

/*
	container.HostConfig:
		Binds           []string      // 此容器的卷绑定列表
		ContainerIDFile string        // 写入containerId的文件（路径）
		LogConfig       LogConfig     // 此容器的日志配置
		NetworkMode     NetworkMode   // 用于容器的网络模式
		PortBindings    nat.PortMap   // 暴露端口（容器）和主机之间的端口映射
		RestartPolicy   RestartPolicy // 用于容器的重新启动策略
		AutoRemove      bool          // 退出时自动删除容器
		VolumeDriver    string        // 用于装载卷的卷驱动程序的名称
		VolumesFrom     []string      // 从其他容器获取的卷列表

		// Applicable to UNIX platforms
		CapAdd          strslice.StrSlice // 要添加到容器的内核功能列表
		CapDrop         strslice.StrSlice // 要从容器中删除的内核功能列表
		CgroupnsMode    CgroupnsMode      // 用于容器的Cgroup命名空间模式
		DNS             []string          `json:"Dns"`        // 要查找的DNS服务器列表
		DNSOptions      []string          `json:"DnsOptions"` // 要查找的DNS选项列表
		DNSSearch       []string          `json:"DnsSearch"`  // 要查找的DNS搜索列表
		ExtraHosts      []string          // 额外主机列表
		GroupAdd        []string          // 容器进程将作为其运行的其他组的列表
		IpcMode         IpcMode           // 用于容器的IPC命名空间
		Cgroup          CgroupSpec        // 用于容器的Cgroup
		Links           []string          // 链接列表（名称：alias form）
		OomScoreAdj     int               // OOM-killing的容器偏好
		PidMode         PidMode           // 用于容器的PID命名空间
		Privileged      bool              // 容器是否处于特权模式
		PublishAllPorts bool              // docker是否应该发布容器的所有暴露端口
		ReadonlyRootfs  bool              // 容器根文件系统是只读的吗
		SecurityOpt     []string          // 用于自定义MLS系统（如SELinux）标签的字符串值列表。
		StorageOpt      map[string]string `json:",omitempty"` // 每个容器的存储驱动程序选项。
		Tmpfs           map[string]string `json:",omitempty"` // 用于集装箱的tmpfs（支架）列表
		UTSMode         UTSMode           // 用于容器的UTS命名空间
		UsernsMode      UsernsMode        // 用于容器的用户命名空间
		ShmSize         int64             // shm内存使用总量
		Sysctls         map[string]string `json:",omitempty"` // 用于容器的命名空间sysctl列表
		Runtime         string            `json:",omitempty"` // 与此容器一起使用的运行时

		// Applicable to Windows
		ConsoleSize [2]uint   // 初始控制台尺寸（高度、宽度）
		Isolation   Isolation // 容器的隔离技术（例如默认、hyperv）

		// 包含容器的资源（cgroup、ulimit）
		Resources

		// 安装容器使用的规格
		Mounts []mount.Mount `json:",omitempty"`

		// MaskedPaths是容器内要屏蔽的路径列表（这将覆盖默认路径集）
		MaskedPaths []string

		// ReadonlyPaths是要在容器内设置为只读的路径列表（这将覆盖默认路径集）
		ReadonlyPaths []string

		// 在容器内运行自定义init，如果为null，则使用守护进程的配置设置
		Init *bool `json:",omitempty"`
*/

/*
	network.NetworkingConfig:
		EndpointsConfig map[string]*EndpointSettings // 为每个连接网络配置端点
*/

/*
	specs.Platform:
		// 架构字段指定CPU架构，例如
		// `amd64` or `ppc64`.
		Architecture string `json:"architecture"`

		// 操作系统指定操作系统，例如“linux”或“windows”。
		OS string `json:"os"`

		// OSVersion是一个可选字段，用于指定操作系统
		// 版本，例如在Windows“10.0.14393.1066”上。
		OSVersion string `json:"os.version,omitempty"`

		// OSFeatures是一个可选字段，用于指定字符串数组，
		// 每个都列出了所需的操作系统功能（例如在Windows“win32k”上）.
		OSFeatures []string `json:"os.features,omitempty"`

		// 变量是一个可选字段，用于指定CPU的变量
		// 示例“v7”用于在体系结构为“arm”时指定ARMv7。
		Variant string `json:"variant,omitempty"`
*/

// ContainerList 获取全部容器列表（含未运行）
func (d *DockerClient) ContainerList() (containerList []types.Container, err error) {
	containerList, err = d.Client.ContainerList(d.Ctx, types.ContainerListOptions{
		All: true,
	})
	return
}

// BatchContainerStop 批量停止容器
func (d *DockerClient) BatchContainerStop(IDMap []string) (success []string, fail []string) {
	for _, cont := range IDMap {
		err := d.Client.ContainerStop(d.Ctx, cont, container.StopOptions{})
		if err != nil {
			fail = append(fail, cont)
		} else {
			success = append(success, cont)
		}
	}
	return
}

// ContainerStop 停止指定容器
func (d *DockerClient) ContainerStop(ID string) error {
	err := d.Client.ContainerStop(d.Ctx, ID, container.StopOptions{})
	return err
}

// ContainerLogs 获取指定容器日志
func (d *DockerClient) ContainerLogs(ID string) (string, error) {
	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := d.Client.ContainerLogs(d.Ctx, ID, options)
	defer func(out io.ReadCloser) {
		err := out.Close()
		if err != nil {
			panic(err)
		}
	}(out)

	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	return buf.String(), err
}

// ContainerCreate 创建容器
func (d *DockerClient) ContainerCreate(
	options *container.Config,
	hostOptions *container.HostConfig,
	networkingOptions *network.NetworkingConfig,
	platform *specs.Platform,
	name string,
) (cont string, err error) {
	resp, err := d.Client.ContainerCreate(d.Ctx, options, hostOptions, networkingOptions, platform, name)
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

// ContainerStart 启动指定容器
func (d *DockerClient) ContainerStart(ID string) error {
	err := d.Client.ContainerStart(d.Ctx, ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}
	return nil
}

// ContainerWait 监听容器状态
func (d *DockerClient) ContainerWait(ID string) (code int, err error) {
	statusCh, errCh := d.Client.ContainerWait(d.Ctx, ID, container.WaitConditionNotRunning)
	select {
	case err = <-errCh:
		if err != nil {
			return 0, err
		}
	case resp := <-statusCh:
		if resp.Error.Message != "" {
			err = errors.New(resp.Error.Message)
		}
		return int(resp.StatusCode), err
	}

	return
}
