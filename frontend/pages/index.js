import { Box, Container } from '@chakra-ui/react'
import Image from 'next/image'

import head from '../public/head.jpg'

const Page = () => {
  return (
    <Image
      width="100%"
      height="100%"
      src={head} />
  )
}

export default Page
