<!--
 * @Author: "xieyuhua" "1130
 * @Date: 2023-12-28 09:06:55
 * @LastEditors: "xieyuhua" "1130
 * @LastEditTime: 2023-12-28 09:10:05
 * @FilePath: \h5\wails-test\README.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# README

## About

This is the official Wails Svelte-TS template. https://wails.io/zh-Hans

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

`wails build`
`go build -tags desktop,production -ldflags "-w -s -H windowsgui"`
