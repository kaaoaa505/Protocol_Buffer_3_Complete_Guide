package org.example;

import com.google.protobuf.InvalidProtocolBufferException;
import dev.local.Dayofweek.DayOfWeek;
import dev.local.Dayofweek.TargetWeekObj;

public class MainTargetWeek {

    public static void main(String[] args) {
        System.out.println("Welcome to target week example...");

        TargetWeekObj.Builder builder = TargetWeekObj.newBuilder();
        builder.setId(123);
        builder.setDayOfWeek(DayOfWeek.SATURDAY);

        TargetWeekObj message  = builder.build();

        byte[] targetBytes = message.toByteArray();
        System.out.println(targetBytes);

        try {
            TargetWeekObj dataFromBinFile = TargetWeekObj.parseFrom(targetBytes);
            System.out.println(dataFromBinFile);
        } catch (InvalidProtocolBufferException e) {
            throw new RuntimeException(e);
        }
    }
}
