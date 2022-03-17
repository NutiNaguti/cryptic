import Head from 'next/head';
import { Box, Container } from '@chakra-ui/react';

const Main = ({ children, router }) => {
  return (
    <Box>
      <Head>
        <meta />
      </Head>
      <Container maxW='container.xl' p={0}>
        {children}
      </Container>
    </Box>
  )
}

export default Main
