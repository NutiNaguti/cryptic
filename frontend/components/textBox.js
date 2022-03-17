import { Box } from '@chakra-ui/react'

const TextBox = ({ children }) => {
  return (
    <Box borderRadius='2xl' bg='gray.100' p='3' color='gray.15' textAlign='center'>
      {children}
    </Box>
  )
}

export default TextBox
