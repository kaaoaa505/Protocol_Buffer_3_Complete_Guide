package org.example;

import dev.local.Nested.*;

public class MainNested {

    public static void main(String[] args) {
        System.out.println("___________________________________________________________");
        System.out.println("------------------- Main Nested: START. -------------------");

        NestedObj message = createNestedObj(34, "Khaled", "Allam");
        System.out.println(message.toByteArray());
        System.out.println(message);

        System.out.println("------------------- Main Nested: END. ---------------------");
        System.out.println("___________________________________________________________");
    }

    public  static  NestedObj createNestedObj(Integer age, String firstName, String lastName){
        NestedObj.Builder  builder = NestedObj.newBuilder();

        NestedObj message = builder.setFirstName(firstName)
                .setLastName(lastName)
                .setAge(age)
                .build();

        return message;
    }
}
