import VMdEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import githubTheme from '@kangc/v-md-editor/lib/theme/github.js';
import '@kangc/v-md-editor/lib/theme/style/github.css';
import enUS from '@kangc/v-md-editor/lib/lang/en-US';

// highlightjs
import hljs from 'highlight.js';

const mdEditor = VMdEditor.use(githubTheme, {
    Hljs: hljs,
    
})

mdEditor.lang.use('en-US', enUS);

export {mdEditor};