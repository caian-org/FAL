namespace Calc
open Amazon.Lambda.Core

[<assembly:LambdaSerializer(typeof<Amazon.Lambda.Serialization.SystemTextJson.DefaultLambdaJsonSerializer>)>]
do ()

[<CLIMutable>]
type Request = { _fal: string }
type ReqInput = { input: string }

type Response = { _fal: string }
type ResOutput = { mutable success: bool; mutable output: string }


module DivideBy =
    open System
    open System.IO
    open System.Text
    open FSharp.Json

    let run(request: Request) =
        printfn "Got event: %A" request

        let res: ResOutput = { success = false; output = null }

        try
            let fal = Json.deserialize<ReqInput> request._fal
            let input = int fal.input

            res.success <- true
            res.output <- string (input / 2)

        with :? System.FormatException as ex ->
            printfn "Could not perform operation: %s" ex.Message

        { _fal = Json.serializeU res }
