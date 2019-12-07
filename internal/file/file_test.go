package file

import (
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("update", false, "update .golden files")

func TestList(t *testing.T) {
	type args struct {
		directory string
	}

	tests := []struct {
		name      string
		args      args
		want      []string
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "directory not found",
			args: args{
				directory: "unknown",
			},
			want:      nil,
			assertion: assert.Error,
		},
		{
			name: "normal",
			args: args{
				directory: "testdata",
			},
			want: []string{"1C-Bitrix",
				"1C",
				"A-Frame",
				"AL",
				"ASPNETCore",
				"ATE",
				"ATS",
				"AWR",
				"Actionscript",
				"Ada",
				"Adobe",
				"AdvancedInstaller",
				"AdventureGameStudio",
				"Agda",
				"AlteraQuartusII",
				"Altium",
				"Android",
				"AndroidStudio",
				"Angular",
				"Anjuta",
				"Ansible",
				"ApacheCordova",
				"ApacheHadoop",
				"AppBuilder",
				"AppCode+all",
				"AppCode+iml",
				"AppCode",
				"AppEngine",
				"AppceleratorTitanium",
				"AptanaStudio",
				"Arcanist",
				"ArchLinuxPackages",
				"Archive",
				"Archives",
				"Assembler",
				"AtmelStudio",
				"Audio",
				"AutomationStudio",
				"Autotools+strict",
				"Autotools",
				"Backup",
				"Ballerina",
				"Basic",
				"Batch",
				"Bazaar",
				"Bazel",
				"BitTorrent",
				"Bitrise",
				"Bitrix",
				"Blackbox",
				"Bloop",
				"Bookdown",
				"Bower",
				"BricxCC",
				"Buck",
				"C++",
				"C",
				"CFWheels",
				"CLion+all",
				"CLion+iml",
				"CLion",
				"CMake",
				"CRBasic",
				"CUDA",
				"CVS",
				"Cake",
				"CakePHP",
				"CakePHP2",
				"CakePHP3",
				"Calabash",
				"Carthage",
				"Ceylon",
				"ChefCookbook",
				"Chocolatey",
				"Clean",
				"Clojure",
				"Cloud9",
				"CocoaPods",
				"Cocos2dx",
				"CocosCreator",
				"Code-Java",
				"Code",
				"CodeBlocks",
				"CodeComposerStudio",
				"CodeIgniter",
				"CodeKit",
				"CodeSniffer",
				"Codeio",
				"CoffeeScript",
				"CommonLisp",
				"Composer",
				"Compressed",
				"CompressedArchive",
				"Compression",
				"Concrete5",
				"Coq",
				"Cordova",
				"CraftCMS",
				"Crashlytics",
				"Crossbar",
				"Crystal",
				"Csharp",
				"D",
				"DBeaver",
				"DM",
				"Dart",
				"DartEditor",
				"Data",
				"DataRecovery",
				"Database",
				"Defold",
				"Delphi",
				"Dframe",
				"Diff",
				"DiskImage",
				"Django",
				"DocFx",
				"Docpress",
				"Docz",
				"DotSettings",
				"DotfilesSh",
				"DotnetCore",
				"Dreamweaver",
				"Dropbox",
				"Drupal",
				"Drupal7",
				"Drupal8",
				"EPiServer",
				"Eagle",
				"Eclipse",
				"EiffelStudio",
				"ElasticBeanstalk",
				"Elisp",
				"Elixir",
				"Elm",
				"Emacs",
				"Ember",
				"Ensime",
				"Erlang",
				"Espresso",
				"Executable",
				"Exercism",
				"ExpressionEngine",
				"ExtJs",
				"Fancy",
				"Finale",
				"Firebase",
				"FlashBuilder",
				"Flask",
				"Flex",
				"FlexBuilder",
				"Flutter",
				"Font",
				"FontForge",
				"ForceDotCom",
				"ForgeGradle",
				"Fortran",
				"FreePascal",
				"FuelPHP",
				"FuseTools",
				"GGTS",
				"GIS",
				"GPG",
				"GWT",
				"Games",
				"Gcov",
				"Genero4GL",
				"Geth",
				"Git",
				"GitBook",
				"Go",
				"Godot",
				"GoodSync",
				"Gradle",
				"Grails",
				"HOL",
				"HSP",
				"Haskell",
				"Helm",
				"Hexo",
				"HomeAssistant",
				"Hugo",
				"HyperledgerComposer",
				"IAR",
				"IAREmbeddedWorkBench",
				"IAR_EWARM",
				"IDAPro",
				"IGORPro",
				"Idris",
				"Images",
				"InforCMS",
				"InforCRM",
				"Intellij+all",
				"Intellij+iml",
				"Intellij",
				"Ionic3",
				"JBoss-4-2-3-GA",
				"JBoss-6-x",
				"JBoss",
				"JBoss4",
				"JBoss6",
				"JDeveloper",
				"JEnv",
				"JGiven",
				"JMeter",
				"JabRef",
				"Java-Web",
				"Java",
				"Jekyll",
				"JetBrains+all",
				"JetBrains+iml",
				"JetBrains",
				"Jigsaw",
				"Joe",
				"Joomla",
				"Julia",
				"JupyterNotebooks",
				"JustCode",
				"KDevelop4",
				"KDiff3",
				"Kate",
				"Keil",
				"Kentico",
				"KiCad",
				"Kirby2",
				"Kobalt",
				"Kohana",
				"KomodoEdit",
				"KonyVisualizer",
				"Kotlin",
				"LAMP",
				"LSspice",
				"LTspice",
				"LaTeX",
				"LabVIEW",
				"LabVIEWNXG",
				"Laravel",
				"Lazarus",
				"Leiningen",
				"LemonStand",
				"Less",
				"LiberoSOC",
				"LibreOffice",
				"Lilypond",
				"Linux",
				"Lithium",
				"Logtalk",
				"Lua",
				"LyX",
				"MATLAB",
				"MEAN",
				"MODX",
				"MPLabX",
				"Magento",
				"Magento1",
				"Magento2",
				"Magic-xpa",
				"Maven",
				"MavensMate",
				"MdBook",
				"Mercurial",
				"Mercury",
				"MetaProgrammingSystem",
				"Metals",
				"Meteor",
				"MeteorJS",
				"MicrosoftOffice",
				"MikroC",
				"Moban",
				"ModelSim",
				"Momentics",
				"MonoDevelop",
				"NCrunch",
				"Nanoc",
				"NativeScript",
				"NesC",
				"NetBeans",
				"Nette",
				"Nikola",
				"Nim",
				"Ninja",
				"Node",
				"NodeChakraTimeTravelDebug",
				"NotepadPP",
				"Nuxt",
				"Nuxtjs",
				"Nwjs",
				"OCaml",
				"OSX",
				"Objective-C",
				"Octave",
				"OctoberCms",
				"Opa",
				"OpenCV",
				"OpenCart",
				"OpenFOAM",
				"OpenFrameworks+VisualStudio",
				"OpenFrameworks",
				"OrCAD",
				"OracleForms",
				"Otto",
				"OxidEshop",
				"PAWN",
				"PHPCodeSniffer",
				"PHPUnit",
				"PSoCCreator",
				"PVS",
				"Packer",
				"Particle",
				"Patch",
				"Perl",
				"Perl6",
				"Phalcon",
				"Phoenix",
				"PhpStorm+all",
				"PhpStorm+iml",
				"PhpStorm",
				"Pimcore",
				"Pimcore4",
				"Pimcore5",
				"PineGrow",
				"PlatformIO",
				"PlayFramework",
				"Plone",
				"Polymer",
				"PowerShell",
				"Prepros",
				"Prestashop",
				"Processing",
				"ProgressABL",
				"PuTTY",
				"Puppet",
				"PureBasic",
				"PureScript",
				"PyCharm+all",
				"PyCharm+iml",
				"PyCharm",
				"Python",
				"QML",
				"Qooxdoo",
				"Qt",
				"QtCreator",
				"R",
				"ROOT",
				"ROS",
				"Racket",
				"Rails",
				"ReactNative",
				"Reasonml",
				"Red",
				"Redcar",
				"Redis",
				"RhodesRhomobile",
				"Rider",
				"Ruby",
				"RubyMine+all",
				"RubyMine+iml",
				"RubyMine",
				"Rust",
				"SAS",
				"SBT",
				"SCons",
				"SSH",
				"SVN",
				"Salesforce",
				"SalesforceDX",
				"Sass",
				"Scala",
				"Scheme",
				"Scrivener",
				"Sdcc",
				"SeamGen",
				"SenchaTouch",
				"Serverless",
				"Shopware",
				"Silverstripe",
				"SketchUp",
				"SlickEdit",
				"Smalltalk",
				"Snap",
				"Snapcraft",
				"Solidity",
				"SolidityTruffle",
				"Sonar",
				"SonarQube",
				"SourcePawn",
				"Spark",
				"Splunk",
				"Spreadsheet",
				"StandardML",
				"Stata",
				"StdLib",
				"Stella",
				"Stellar",
				"Stylus",
				"SublimeText",
				"SugarCRM",
				"Swift",
				"SwiftPM",
				"SwiftPackageManager",
				"Symfony",
				"SymphonyCMS",
				"Synology",
				"SynopsysVCS",
				"THEOS-Tweak",
				"TYPO3-composer",
				"Tags",
				"TarmaInstallMate",
				"TeX",
				"Terraform",
				"Terragrunt",
				"Test",
				"TestComplete",
				"Testinfra",
				"Text",
				"TextMate",
				"Textpattern",
				"ThinkPHP",
				"TortoiseGit",
				"Tower",
				"TurboGears2",
				"TwinCAT",
				"Typings",
				"Typo3",
				"Umbraco",
				"Unity",
				"UnrealEngine",
				"VLab",
				"VVVV",
				"Vaadin",
				"Vagrant",
				"Valgrind",
				"Vapor",
				"Vertx",
				"Video",
				"Vim",
				"VirtualEnv",
				"Virtuoso",
				"VisualStudio",
				"VisualStudioCode",
				"Vivado",
				"Vue",
				"Vuejs",
				"Waf",
				"Wakanda",
				"Web",
				"WebMethods",
				"WebStorm+all",
				"WebStorm+iml",
				"WebStorm",
				"WerckerCLI",
				"Windows",
				"Wintersmith",
				"WordPress",
				"Wyam",
				"XText",
				"XamarinStudio",
				"Xcode",
				"XcodeInjection",
				"Xilinx",
				"XilinxISE",
				"XilinxVivado",
				"Xill",
				"Xojo",
				"Y86",
				"Yeoman",
				"Yii",
				"Yii2",
				"ZendFramework",
				"Zephir",
				"Zsh",
				"ZukenCR8000",
				"baserCMS",
				"bluej",
				"certificates",
				"direnv",
				"dotenv",
				"e2studio",
				"easybook",
				"fastlane",
				"floobits",
				"fsharp",
				"greenfoot",
				"grunt",
				"infer",
				"jspm",
				"librarian-chef",
				"m2e",
				"macOS",
				"mule",
				"oXygenXMLEditor",
				"pH7CMS",
				"pico8",
				"premake-gmake",
				"puppet-librarian",
				"pydev",
				"react",
				"venv",
				"zig"},
			assertion: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.directory)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGenerate(t *testing.T) {
	type args struct {
		directory string
		items     []string
	}

	tests := []struct {
		name      string
		args      args
		wantW     string
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "only ignore",
			args: args{
				directory: "testdata",
				items:     []string{"ELM"},
			},
			wantW:     "only-ignore.golden",
			assertion: assert.NoError,
		},
		{
			name: "only ignore (duplicated items)",
			args: args{
				directory: "testdata",
				items:     []string{"ELM", "elm"},
			},
			wantW:     "only-ignore.golden",
			assertion: assert.NoError,
		},
		{
			name: "with patch",
			args: args{
				directory: "testdata",
				items:     []string{"go", "elm"},
			},
			wantW:     "with-patch.golden",
			assertion: assert.NoError,
		},
		{
			name: "with stack",
			args: args{
				directory: "testdata",
				items:     []string{"lamP"},
			},
			wantW:     "with-stack.golden",
			assertion: assert.NoError,
		},
		{
			name: "with duplicated lines",
			args: args{
				directory: "testdata",
				items:     []string{"go", "c"},
			},
			wantW:     "with-duplicated-lines.golden",
			assertion: assert.NoError,
		},
		{
			name: "with undefined",
			args: args{
				directory: "testdata",
				items:     []string{"go", "go++"},
			},
			wantW:     "with-undefined.golden",
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := Generate(w, `testdata`, tt.args.items...)
			tt.assertion(t, err)

			goldenPath := filepath.Join(`_golden`, tt.wantW)

			if *update {
				require.NoError(t, ioutil.WriteFile(goldenPath, w.Bytes(), 0644))
			}

			expected, err := ioutil.ReadFile(goldenPath)
			require.NoError(t, err)
			assert.Equal(t, expected, w.Bytes())
		})
	}
}

func TestGenerate_UnknownDirectory(t *testing.T) {
	w := &bytes.Buffer{}
	assert.Error(t, Generate(w, `unknown`))
}
