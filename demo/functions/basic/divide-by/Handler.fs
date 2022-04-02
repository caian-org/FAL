namespace Calc

open Amazon.Lambda.Core

[<assembly:LambdaSerializer(typeof<Amazon.Lambda.Serialization.SystemTextJson.DefaultLambdaJsonSerializer>)>]
do ()

[<CLIMutable>]
type Request = { _fal : string }

type KeyNotFound = System.Collections.Generic.KeyNotFoundException


module DivideBy =
    open System
    open System.IO
    open System.Text

    let run(req : Request) =
        printfn "Got event: %A" req

        if (req._fal = null)
            then raise (KeyNotFound("Event is missing \"_fal\" property"))

        let input = int req._fal
        let output = string(input / 2)

        { _fal = output }
