package com.plusone;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.annotation.PropertyAccessor;
import com.fasterxml.jackson.annotation.JsonAutoDetect.Visibility;

public class Mapper
{
    public static ObjectMapper build()
    {
        return new ObjectMapper().setVisibility(PropertyAccessor.FIELD, Visibility.ANY);
    }
}
