import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import React, { useState } from "react";
import { IoAdd } from "react-icons/io5";
import { BASE_URL } from "../App";

const TodoForm = () => {
  const [newTodo, setNewTodo] = useState("");
  const queryClient = useQueryClient();

  const { mutate: createTodo, isPending: isCreating } = useMutation({
    mutationKey: ["createTodo"],
    mutationFn: async (e: React.FormEvent) => {
      e.preventDefault();
      try {
        const res = await fetch(BASE_URL + `/todo`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ body: newTodo }),
        });

        const data = await res.json();

        if (!res.ok) {
          throw new Error(data.error || "Something went wrong");
        }

        setNewTodo("");
        return data;
      } catch (error: any) {
        throw new Error(error);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: (error: any) => {
      alert(error.message);
    },
  });

  return (
    <form onSubmit={createTodo}>
      <Flex gap={2}>
        <Input
          type="text"
          borderWidth={3}
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
          ref={(input) => input && input.focus()}
          placeholder="Add todo"
        />
        <Button
          mx={2}
          type="submit"
          _active={{ transform: "scale(.97)" }}
          colorScheme="blue"
        >
          {isCreating ? <Spinner size={"xs"} /> : <IoAdd size={30} />}
        </Button>
      </Flex>
    </form>
  );
};

export default TodoForm;
