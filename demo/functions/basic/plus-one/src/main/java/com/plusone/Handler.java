package com.plusone;

import java.util.Collections;
import java.util.Map;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;

public class Handler implements RequestHandler<Map<String, String>, Response>
{
    @Override
    public Response handleRequest (Map<String, String> event, Context context)
    {
        System.out.println("Got event: " + event);

        String falKey = "_fal";
        if (!event.containsKey(falKey)) {
            throw new NullPointerException(String.format("Event is missing \"%s\" property", falKey));
        }

        int input = Integer.parseInt(event.get(falKey));
        return new Response(String.valueOf(input + 1));
    }
}
