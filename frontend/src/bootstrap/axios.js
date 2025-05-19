import axios from 'axios';

const setupAxios = () => {
    window.axios = axios;
    window.axios.defaults.headers.common['Telegram-Init'] = window.Telegram.WebApp.initData;
}

export {setupAxios}