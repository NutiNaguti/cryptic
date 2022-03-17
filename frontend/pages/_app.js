import { ChakraProvider } from '@chakra-ui/react'
import Layout from '../components/layouts/main'

const WebSite = ({ Component, pageProps, router }) => {
  return (
    <ChakraProvider>
      <Component {...pageProps} key={router.route} />
    </ChakraProvider>
  )
}

export default WebSite
