import { extendTheme } from '@chakra-ui/react'
import { theme as chakraTheme } from '@chakra-ui/react'
import { createBreakpoints } from '@chakra-ui/react'

const fonts = {
  ...chakraTheme.fonts,
  body: ``
}

const overrides = {
  ...chakraTheme,
  fonts
}

const customTheme = extendTheme({ overrides })

export default customTheme
