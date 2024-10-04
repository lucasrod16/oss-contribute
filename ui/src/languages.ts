export const languageImageMap = {
	// core languages
	C: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/c/c-original.svg",
	"C++": "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/cplusplus/cplusplus-original.svg",
	"C#": "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/csharp/csharp-original.svg",
	Go: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg",
	Java: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/java/java-original.svg",
	JavaScript: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/javascript/javascript-original.svg",
	PHP: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/php/php-original.svg",
	Python: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/python/python-original.svg",
	Ruby: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/ruby/ruby-original.svg",
	Rust: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/rust/rust-original.svg",
	Scala: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/scala/scala-original.svg",
	Swift: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/swift/swift-original.svg",
	TypeScript: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/typescript/typescript-original.svg",

	// others
	Dart: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/dart/dart-original.svg",
	Svelte: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/svelte/svelte-original.svg",
	"Jupyter Notebook": "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/jupyter/jupyter-original.svg",
	Crystal: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/crystal/crystal-original.svg",
	Lua: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/lua/lua-original.svg",
	Kotlin: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/kotlin/kotlin-original.svg",
	Jsonnet: "https://www.svgrepo.com/show/373716/jsonnet.svg",
	Markdown: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/markdown/markdown-original.svg",
	OCaml: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/ocaml/ocaml-original.svg",
	"Emacs Lisp": "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/emacs/emacs-original.svg",
	"F#": "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/fsharp/fsharp-original.svg",
	HTML: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/html5/html5-original.svg",
	SystemVerilog: "https://www.svgrepo.com/show/374115/systemverilog.svg",
	Shell: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/bash/bash-original.svg",
	Roff: "https://repository-images.githubusercontent.com/564045681/5a1306b8-3695-40f8-ac0b-8da2e9ddfbe4",
	Raku: "https://upload.wikimedia.org/wikipedia/commons/8/85/Camelia.svg",
	Julia: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/julia/julia-original.svg",
	CoffeeScript: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/coffeescript/coffeescript-original.svg",

	// unknown/empty
	"": "https://www.svgrepo.com/show/397695/red-question-mark.svg",
} as const;

export type Language = keyof typeof languageImageMap;

export function getLanguageImageURL(language: Language): string {
	return languageImageMap[language];
}
