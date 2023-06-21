package org.example;

import com.host.local.options.OptionsObj;
import com.host.local.options.OptionsV2Obj;
import dev.local.Nested;

public class MainOptions {
    public static void main(String[] args) {
        System.out.println("___________________________________________________________");
        System.out.println("------------------- Main Options: START. -------------------");

        OptionsObj.Builder builder = OptionsObj.newBuilder();
        OptionsObj message = builder.setId(123).build();

        OptionsV2Obj.Builder builderV2 = OptionsV2Obj.newBuilder();
        OptionsV2Obj messageV2 = builderV2.setUuid("6d02da22-3137-46d5-b395-f2b5b8c94b80").build();

        System.out.println(message);
        System.out.println(messageV2);

        System.out.println("------------------- Main Options: END. ---------------------");
        System.out.println("___________________________________________________________");
    }
}
