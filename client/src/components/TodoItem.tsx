import { Badge, Box, Flex, Spinner, Text } from "@chakra-ui/react"
import { FaCheckCircle } from "react-icons/fa"
import { MdDelete } from "react-icons/md"
import { Todo } from "./TodoList" // Pastikan Anda memiliki tipe Todo yang tepat
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { BASE_URL } from "../App"

const TodoItem = ({ todo }: { todo: Todo }) => {
  const queryClient = useQueryClient()

  const { mutate: updateTodo, isPending: isUpdating } = useMutation({
    mutationKey: ["updateTodo", todo.id], // Gunakan todo.id
    mutationFn: async () => {
      if (todo.completed) {
        alert("Todo is already completed")
        return // Hentikan eksekusi jika todo sudah selesai
      }

      try {
        const res = await fetch(`${BASE_URL}/todos/${todo.id}`, {
          method: "PATCH",
        })
        const data = await res.json()
        if (!res.ok) {
          throw new Error(data.error || "Something went wrong")
        }
        return data
      } catch (error) {
        console.error(error)
        alert(error)
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] })
    },
  })

  const { mutate: deleteTodo, isPending: isDeleting } = useMutation({
    mutationKey: ["deleteTodo", todo.id], // Gunakan todo.id
    mutationFn: async () => {
      try {
        const res = await fetch(`${BASE_URL}/todos/${todo.id}`, {
          method: "DELETE",
        })
        const data = await res.json()
        if (!res.ok) {
          throw new Error(data.error || "Something went wrong")
        }
        return data
      } catch (error) {
        console.error(error)
        alert(error)
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] })
    },
  })

  return (
    <Flex gap={2} alignItems={"center"}>
      <Flex
        flex={1}
        alignItems={"center"}
        border={"1px"}
        borderColor={"gray.600"}
        p={2}
        borderRadius={"lg"}
        justifyContent={"space-between"}>
        <Text
          color={todo.completed ? "green.200" : "yellow.100"}
          textDecoration={todo.completed ? "line-through" : "none"}>
          {todo.body}
        </Text>
        {todo.completed ? (
          <Badge ml="1" colorScheme="green">
            Done
          </Badge>
        ) : (
          <Badge ml="1" colorScheme="yellow">
            In Progress
          </Badge>
        )}
      </Flex>
      <Flex gap={2} alignItems={"center"}>
        <Box
          color={"green.500"}
          cursor={"pointer"}
          onClick={() => updateTodo()}>
          {isUpdating ? <Spinner size={"sm"} /> : <FaCheckCircle size={20} />}
        </Box>
        <Box color={"red.500"} cursor={"pointer"} onClick={() => deleteTodo()}>
          {isDeleting ? <Spinner size={"sm"} /> : <MdDelete size={25} />}
        </Box>
      </Flex>
    </Flex>
  )
}

export default TodoItem
