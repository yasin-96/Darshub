const gitModule = {
    namespaced: true,
    state: {
      gitVersion: process.env.VUE_APP_GIT_VERSION,
      gitHash: process.env.VUE_APP_GIT_HASH,
      gitDateTimeOfLastCommit: process.env.VUE_APP_GIT_LAST_COMMIT_DATE,
    },
    getters: {
      /**
       * Returns the date in form YYYY.MM.DD from last commit
       * @param {Object} state The local state of module
       * @returns {String} The date of last commit
       */
      onlyDate: function(state) {
        return String(state.gitDateTimeOfLastCommit)
          .trim()
          .split('T')[0];
      },
  
      /**
       * Returns from last commit a short version of the hash id.
       * @param {Object} state The local state of module
       * @returns {String}
       *  The HashCode of last commit in short version or
       *  null if webpack could not set the enviroment variables.
       */
      minimalHashCode(state) {
        if (state.gitHash) {
          return state.gitHash.substring(0, 7);
        }
        return null;
      },
  
      /**
       * Returns a tag or version number if this was set in the last commit.
       * @param {Object} state The local state of module
       * @returns {String}
       *  The TAG of last commit or null if webpack could not set the enviroment variables
       */
      checkVersionTag(state) {
        if (state.gitHash.substring(0, 7) != state.gitVersion) {
          return state.gitVersion;
        }
        return null;
      },
    },
  };
  
  export default gitModule;