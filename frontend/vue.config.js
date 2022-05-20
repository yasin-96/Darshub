const { GitRevisionPlugin } = require('git-revision-webpack-plugin');

module.exports = {
  transpileDependencies: ["vuetify"],
  productionSourceMap: false, // Do not create js source maps when building the App
  devServer: {
    https: false,
    port: 88
  },
  chainWebpack: (config) => {
    config
    .plugin('define')
    .tap((args) => {
      const gitRevisionPlugin = new GitRevisionPlugin({
        commithashCommand: 'rev-list --max-count=1 --no-merges HEAD',
        versionCommand: 'describe --always --tags',
      });

      args[0]['process.env']['VUE_APP_GIT_VERSION'] = JSON.stringify(gitRevisionPlugin.version());
      args[0]['process.env']['VUE_APP_GIT_HASH'] = JSON.stringify(gitRevisionPlugin.commithash());
      args[0]['process.env']['VUE_APP_GIT_LAST_COMMIT_DATE'] = JSON.stringify(gitRevisionPlugin.lastcommitdatetime());
      return args;
    });

    config.module.rule('pdf')
      .test(/\.pdf$/)
      .use('file-loader').loader('file-loader');
  },
};
