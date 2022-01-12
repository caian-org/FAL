package com.plusone;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

public class Response
{
    public String _fal;

    public Response(Object data)
    {
        try {
            this._fal = Mapper.build().writeValueAsString(data);
        } catch (JsonProcessingException e) {
            throw new RuntimeException("Could not process JSON: " + e.getMessage());
        }
    }

    public static Builder builder()
    {
        return new Builder();
    }

    public static class Builder
    {
        private boolean success;
        private String output;

        public Builder setSuccess(boolean success)
        {
            this.success = success;
            return this;
        }

        public Builder setOutput(String output)
        {
            this.output = output;
            return this;
        }

        public Response build()
        {
            return new Response(this);
        }
    }
}
