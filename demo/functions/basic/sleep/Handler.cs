using System;
using System.Threading;

[assembly:LambdaSerializer(typeof(Amazon.Lambda.Serialization.SystemTextJson.DefaultLambdaJsonSerializer))]

namespace Basic;

public class Sleep
{
   public void Run(Request req)
   {
       var seconds = req._fal is string && req._fal.Length > 0
           ? Convert.ToInt32(req._fal)
           : 5;

       Console.WriteLine("Seconds to sleep: " + seconds);

       Thread.Sleep(seconds * 1000);
       Console.WriteLine("Finished");
   }
}

public class Request
{
    public string _fal {get; set;}
}
