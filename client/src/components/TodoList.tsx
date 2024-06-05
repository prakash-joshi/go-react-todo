import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import { useState } from "react";
import TodoItem from "./TodoItem";

const TodoList = () => {
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const todos = [
    { _id: 1, completed: true, body: "Do laundry" },
    { _id: 2, completed: true, body: "Cook dinner" },
    { _id: 3, completed: false, body: "Walk the dog" },
    { _id: 4, completed: false, body: "buy groceries" },
  ];

  return (
    <>
      <Text
        fontSize={"4xl"}
        textTransform={"uppercase"}
        fontWeight={"bold"}
        textAlign={"center"}
        my={2}
      >
        Today's Task
      </Text>
      {isLoading && (
        <Flex justifyContent={"center"} my={4}>
          <Spinner size={"xl"} />
        </Flex>
      )}
      {!isLoading && todos?.length === 0 && (
        <Stack alignItems={"center"} gap={3}>
          <Text fontSize={"xl"} textAlign={"center"} color={"grey.500"}>
            All tasks completed! 🤞
          </Text>
          <img src="./go.png" alt="Go logo" width={70} height={70} />
        </Stack>
      )}
      <Stack gap={3}>
        {todos?.map((todo) => (
          <TodoItem key={todo?._id} todo={todo} />
        ))}
      </Stack>
    </>
  );
};

export default TodoList;
