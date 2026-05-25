Here are some rough ideas on how the content-composer engine should work based on what I know so far.

# Major Components:
- Everything design is defined and driven by EditorJS. 
- Everything relating to rendering these designs are defined by Web.Site. 
- The middle ground between both grounds is the actual engine which leverages Golang  tmplFunc
-


# TODO:

## Content-Composer: Website Editor and Site builder

### Content-Composer Engine:
- [] Implement the initial version of the compiler that parses and compiles between the editor JS data provided and the content composer engine's matching template definitions for EditorJS tools (Block tools)
- [] Add examples and tests for further documentation of the usage 
  - [] Layout system, page definitions and actions 
  - [] page definitions and actions (how to structure code to define a specific page, if making custom, or how to use the editor to define one)
  - [] Site building (how to structure code to define a specific site, if making custom, or how to use the editor to define one)
  - [] Base tools (custom imlementations of Text and typography, list tools, Media and Embedded, Table, Chart, Button (with custom action setter), Forms)
  - [] Components definitions (combining and integrating base tools to get a custom plugin), 
  - [] Themes (how to structure code to put related plugins together))
- [] Provide a spec on how the plugin system works with the former rule


### Content-Composer Worksapce:
- [] Make a nice looking interface integrating all the usages defined earlier
- [] Create components and themes for a sample new and old aiesec.ca page (and site) 
- [] Provide a comprehensive tutorial on how to build such a page and leverage tools for marketing

## Auth Manager: Identity management for access to aie.tools' services
# Role Management: 
// Go through the documentation aiesec IM to see what exactly the roles are in aiesec so as to set custom platform roles for Content composer that allows certain roles/LCs to perform actions on sites/pages based on base and custome rules set by super admin.
// Define basic admin roles for aie.tools that pull from actual aiesec roles. (most likely MCP and MCVP IM as super admin + other roles as MC may see fit but also based on how other organizations set roles) 
// For the sake of simplicity of showcase, we are only going to heavily work in point 1 and the super admin role in point 2
