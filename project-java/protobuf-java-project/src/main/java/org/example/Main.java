package org.example;

import dev.sample.local.Sample.SampleObj;

import java.io.FileOutputStream;
import java.io.IOException;
import java.util.Arrays;

public class Main {
    public static void main(String[] args) {
        System.out.println("Hello world!");

        SampleObj.Builder builder = SampleObj.newBuilder();
        builder.setId(123)
                .setName("Khaled Allam")
                .setIsValid(true)
                .addAllStrList(Arrays.asList("Khaled", "Allam", "Ahmed", "Omar", "Abdalla", "Allam"));

        System.out.println(builder.getName());
        System.out.println(builder.getStrList(2));
        System.out.println(builder.getStrList(3));
        System.out.println(builder.getStrList(4));

        SampleObj message = builder.build();

        try {
            FileOutputStream outputStream = new FileOutputStream("sample_output.bin");
            message.writeTo(outputStream);
            outputStream.close();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

        byte[] targetBytes = message.toByteArray();
        System.out.println(targetBytes);
    }
}