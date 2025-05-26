import data from './webAppData.json'

const setupTelegram = () => {
    if (window.Telegram.WebApp.initData) {
        return
    }

    window.Telegram = data
}

export { setupTelegram }