export default {
  plugins: [
    {
      name: 'preset-default',
      params: {
        overrides: {
          // Keep viewBox for responsiveness
          removeViewBox: false,
        },
      },
    },
    // Remove extra whitespace by recalculating viewBox
    'removeDimensions',
    'removeUselessStrokeAndFill',
  ],
};
