package com.plusone;

import java.util.Collections;
import java.util.Map;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
// import com.fasterxml.jackson.databind.ObjectMapper;

public class Handler implements RequestHandler<Map<String, String>, Response>
{
    private Response sendResponse (String output)
    {
        return Response.builder()
            .setSuccess(output != null)
            .setOutput(output == null ? "" : output)
            .build();
    }

    @Override
    public Response handleRequest (Map<String, String> event, Context context)
    {
        System.out.println("Got event: " + event);
        String falKey = "_fal";

        try {
            if (!event.containsKey(falKey)) {
                throw new NullPointerException(String.format("Event is missing \"%s\" property", falKey));
            }

            Request fal = Mapper.build().readValue(event.get(falKey), Request.class);
            int input = Integer.parseInt(fal.getInput());

            return this.sendResponse(String.valueOf(input + 1));
        } catch (Exception e) {
            System.out.println("Could not perform operation: " + e.getMessage());
        }

        return this.sendResponse(null);
    }
}
