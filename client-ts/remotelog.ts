/* tslint:disable */
// remotelog v1.0.0
// --
// This file has been generated by https://github.com/webrpc/webrpc using gen/typescript
// Do not edit by hand. Update your webrpc schema and re-generate.


export interface RemoteLog {
  ping(headers?: object): Promise<PingReturn>
  log(args: LogArgs, headers?: object): Promise<LogReturn>
}

export interface PingArgs {
}

export interface PingReturn {
  status: boolean
  ts: string  
}
export interface LogArgs {
  msg: string
  object?: {[key: string]: any}
}

export interface LogReturn {
  status: boolean  
}


  
// Client
export class RemoteLog implements RemoteLog {
  private hostname: string
  private fetch: Fetch
  private path = '/rpc/RemoteLog/'

  constructor(hostname: string, fetch: Fetch) {
    this.hostname = hostname
    this.fetch = fetch
  }

  private url(name: string): string {
    return this.hostname + this.path + name
  }
  
  ping = (headers?: object): Promise<PingReturn> => {
    return this.fetch(
      this.url('Ping'),
      createHTTPRequest({}, headers)
      ).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          status: <boolean>(_data.status), 
          ts: <string>(_data.ts)
        }
      })
    })
  }
  
  log = (args: LogArgs, headers?: object): Promise<LogReturn> => {
    return this.fetch(
      this.url('Log'),
      createHTTPRequest(args, headers)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          status: <boolean>(_data.status)
        }
      })
    })
  }
  
}

  
export interface WebRPCError extends Error {
  code: string
  msg: string
	status: number
}

const createHTTPRequest = (body: object = {}, headers: object = {}): object => {
  return {
    method: 'POST',
    headers: { ...headers, 'Content-Type': 'application/json' },
    body: JSON.stringify(body || {})
  }
}

const buildResponse = (res: Response): Promise<any> => {
  return res.text().then(text => {
    let data
    try {
      data = JSON.parse(text)
    } catch(err) {
      throw { code: 'unknown', msg: `expecting JSON, got: ${text}`, status: res.status } as WebRPCError
    }
    if (!res.ok) {
      throw data // webrpc error response
    }
    return data
  })
}

export type Fetch = (input: RequestInfo, init?: RequestInit) => Promise<Response>
