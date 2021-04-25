module.exports = {
	base: '/',//基础路径
	head: [
		['meta', { name: 'keywords', content: 'vuepress 模板' }]
	],
	host: '0.0.0.0',
	dest: './dist',//打包输出目录
	markdown: {
		lineNumbers: true
	},
	plugins: ['@vuepress/back-to-top'],
	locales: {
		// 键名是该语言所属的子路径
		// 作为特例，默认语言可以使用 '/' 作为其路径。
		'/': {
			lang: 'en-US', // 将会被设置为 <html> 的 lang 属性
			title: 'Tigago Web Framework',
			description: `Tigago is a modular, high performance, production level go basic development framework`
		},
		'/zh/': {
			lang: 'zh-CN',
			title: 'Tigago Web Framework',
			description: 'Tigago 是一款模块化、高性能、生产级的Go基础开发框架'
		}
	},
	theme: undefined,
	themeConfig: {
		logo: '/images/logo.jpg',
		repo: 'https://github.com/tigateam/tigago', // 顶部显示的github地址
		docsRepo: 'https://github.com/tigateam/tigago-website', //文档内改进页面的仓库地址
		docsDir: 'docs', //文档目录
		docsBranch: 'master', //分支
		editLinks: true,
		locales: {
			'/': {
				selectText: 'Languages',
				label: 'English',
				repoLabel: 'Source',
				editLinkText: 'Edit this page on GitHub',
				lastUpdated: 'Last Updated',
				serviceWorker: {
					updatePopup: {
						message: "New content is available.",
						buttonText: "Refresh"
					}
				},
				algolia: {},
				sidebarDepth: 1,
				sidebar: 'auto',
				nav: [
					{
						text: 'Home', link: '/'
					},
					{
						text: 'Documentation', link: '/documentation/Introduction'
					},

				],
				sidebar: {
					'/documentation/': [
						{
							title: 'Introduction',
							path: '/documentation/Introduction',
						}, {
							title: 'JoinUs',
							path: '/documentation/JoinUs',
						}, {
							title: 'Preparation',
							path: '/documentation/Preparation/',
							collapsable: true,
							children: [
								{
									title: 'EnvironmentInstallation',
									path: '/documentation/Preparation/EnvironmentInstallation',
								}
							]
						},
						{
							title: 'UserMustRead',
							collapsable: true,
							children: [
								{
									title: 'Development specification',
									path: '/documentation/UserMustRead/DevelopmentSpecification',
								}
							]
						},
						{
							title: 'QuickStart',
							path: '/documentation/QuickStart/',
							collapsable: true,
							children: [
							]
						},
						{
							title: 'ProjectDeployment',
							collapsable: true,
							children: [
								{
									title: "Nginx",
									path: '/documentation/ProjectDeployment/Nginx'
								},
								{
									title: "Docker",
									path: '/documentation/ProjectDeployment/Docker'
								}
							]
						},
						{
							title: 'Web Service Development',
							collapsable: true,
							children: [
							]
						},
						{
							title: 'Modules',
							collapsable: true,
							children: [
							]
						},
						{
							title: 'Framework Learning Tutorial',
							collapsable: true,
							children: [
								{
									title: 'The Official Tutorial',
									collapsable: true,
									children: [
									]
								}, {
									title: 'The Community tutorial',
									collapsable: true,
									children: [
									]
								},
							]
						},
						{
							title: 'Framework User Case',
							collapsable: true,
							children: [
							]
						},
						{
							title: 'Change Log',
							path: '/documentation/ChangeLog/',
							collapsable: true,
							children: [
								{
									title: "V1.0.0",
									path: '/documentation/ChangeLog/V1.0.0',
								},
							]
						},
						{
							title: 'FAQ List',
							path: '/documentation/FAQ',
						},
					],
				}
			},
			'/zh/': {
				// 多语言下拉菜单的标题
				selectText: '选择语言',
				// 该语言在下拉菜单中的标签
				label: '简体中文',

				repoLabel: '查看源码',
				// 以下为可选的编辑链接选项
				// 默认为 "Edit this page"
				editLinkText: '帮助我们改善此页面！',
				// 最后更新时间
				lastUpdated: '最后更新时间',
				// Service Worker 的配置
				serviceWorker: {
					updatePopup: {
						message: "发现新内容可用.",
						buttonText: "刷新"
					}
				},
				// 当前 locale 的 algolia docsearch 选项
				algolia: {},
				sidebarDepth: 1,
				sidebar: 'auto',
				nav: [
					{
						text: '主页', link: '/zh/'
					},
					{
						text: '官方文档', link: '/zh/documentation/Introduction'

					},
				],
				sidebar: {
					'/zh/documentation/': [
						{
							title: '框架介绍',
							path: '/zh/documentation/Introduction',
							sidebarDepth: 0
						}, {
							title: '加入我们',
							path: '/zh/documentation/JoinUs',
							sidebarDepth: 0
						}, {
							title: '准备工作',
							path: '/zh/documentation/Preparation/',
							collapsable: true,
							children: [
								{
									title: "环境安装",
									path: '/zh/documentation/Preparation/EnvironmentInstallation',
								}
							]
						},
						{
							title: '用户必读',
							collapsable: true,
							children: [
								{
									title: "开发规范",
									path: '/zh/documentation/UserMustRead/DevelopmentSpecification',
								}
							]
						}, {
							title: '快速开始',
							path: '/zh/documentation/QuickStart/',
							collapsable: true,
							children: [
								// '/zh/documentation/QuickStart/',
							]
						}, {
							title: '项目部署',
							collapsable: true,
							children: [
								{
									title: "Nginx",
									path: '/zh/documentation/ProjectDeployment/Nginx'
								},
								{
									title: "Docker",
									path: '/zh/documentation/ProjectDeployment/Docker'
								}
							]
						},
						{
							title: 'Web服务开发',
							collapsable: true,
							children: [
							]
						},
						{
							title: 'Modules',
							collapsable: true,
							children: [
							]
						}, {
							title: '框架学习教程',
							collapsable: true,
							children: [
								{
									title: '官方教程',
									collapsable: true,
									children: [
									]
								}, {
									title: '社区教程',
									collapsable: true,
									children: [
									]
								},
							]
						},
						{
							title: '框架用户案例',
							collapsable: true,
							children: [
							]
						},
						{
							title: 'Change Log',
							path: '/zh/documentation/ChangeLog/',
							collapsable: true,
							children: [
								{
									title: "V1.0.0",
									path: '/zh/documentation/ChangeLog/V1.0.0',
								},
							]
						},
						{
							title: 'FAQ 疑问解答',
							path: '/zh/documentation/FAQ',
						},
					]
				}
			}
		}
	},
}