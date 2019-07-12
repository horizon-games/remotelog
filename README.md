remotelog
=========

a remote console logger for web apps


## Usage

*Server*:
1. Install the server, `go get -u github.com/horizon-games/remotelog`
2. Run the server `remotelog` (default port is 1111), or run `remotelog -port=1111`

*Client integration*:
Simply copy either `client-js/remotelog.js` or `client-ts/remotelog.ts` to your project
and then make calls like so..

```
  const svcFetch = window.fetch.bind(window)
  let remotelog = new RemoteLog('http://127.0.0.1:1111', svcFetch)
  window.remotelog = remotelog

  // later in your app
  window.remotelog.log({ msg: 'just a regular console log' })
  window.remotelog.log({ msg: 'log me an object too', object: JSON.stringify(anObject) })
```
