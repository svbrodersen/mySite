// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func BaseLayout() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link rel=\"stylesheet\" href=\"/static/css/style.css\"><title>Personal Portfolio</title><script src=\"/static/javascript/htmx.min.js\"></script></head><body><div class=\"flex flex-col h-screen\"><header class=\"flex-shrink\"><nav><ul class=\"w-screen flex justify-end p-4 pr-32 space-x-3\"><li class=\"nav-item\"><a hx-replace-url=\"true\" hx-patch=\"/about\" hx-trigger=\"click\" hx-target=\"#contents\">About Me</a></li><li class=\"nav-item animate-delay-75\"><a hx-replace-url=\"true\" hx-patch=\"/projects\" hx-trigger=\"click\" hx-target=\"#contents\">Projects</a></li><li class=\"nav-item animate-delay-100\"><a hx-patch=\"/cv\" hx-replace-url=\"true\" hx-trigger=\"click\" hx-target=\"#contents\">CV</a></li><li class=\"nav-item animate-delay-150\"><a hx-patch=\"/contact\" hx-replace-url=\"true\" hx-trigger=\"click\" hx-target=\"#contents\">Contact</a></li></ul></nav></header><div class=\"flex-grow h-auto mx-auto justify-start items-start fill\n\t\tbasis-full w-1/2\" id=\"contents\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><footer class=\"flex flex-shrink items-end\"><ul class=\"pb-4 w-screen flex justify-center space-x-2 justify-items-end\"><li class=\"pt-1\"><a href=\"https://www.linkedin.com/in/simon-vinding-brodersen/\"><img width=\"24\" height=\"24\" src=\"/static/img/LinkedIn-Logos/LI-In-Bug.png\" alt=\"LinkedIn SVG\"></a></li><li><a href=\"https://github.com/svbrodersen\"><picture><source srcset=\"/static/img/github-mark/github-mark-white.png\" media=\"(prefers-color-scheme: dark)\"> <img width=\"24\" height=\"24\" src=\"/static/img/github-mark/github-mark.png\" alt=\"Github SVG\"></picture></a></li></ul></footer></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
