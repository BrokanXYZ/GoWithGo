import { createMuiTheme } from '@material-ui/core/styles';

const theme = createMuiTheme({
    spacing: factor => `${0.25 * factor}rem`, // (Bootstrap strategy)
    palette: {
      primary: {
        main: '#eeeeee',
        //light: '#eeffff',
        //dark: '#8aacc8',
        // contrastText: will be calculated to contrast with palette.primary.main
      },
      secondary: {
        main: '#78909c',
        //light: '#ffd449',
        //dark: '#c67400',
        // contrastText: will be calculated to contrast with palette.primary.main
      },
      // Used by `getContrastText()` to maximize the contrast between
      // the background and the text.
      contrastThreshold: 3,
      // Used by the functions below to shift a color's luminance by approximately
      // two indexes within its tonal palette.
      // E.g., shift from Red 500 to Red 300 or Red 700.
      tonalOffset: 0.2,
    },
});

export default theme;