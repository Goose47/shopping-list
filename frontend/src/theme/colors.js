
const colors = {
    'telegram': {},
    'jackass': {
        'accent_text_color': '#9D9101',
        'bg_color': '#9E9764',
        'bottom_bar_bg_color': '#474B4E',
        'button_color': '#2D572C',
        'button_text_color': '#F44611',
        'destructive_text_color': '#79553D',
        'header_bg_color': '#C7B446',
        'hint_color': '#E63244',
        'link_color': '#5D9B9B',
        'secondary_bg_color': '#1E5945',
        'section_bg_color': '#6A5D4D',
        'section_header_text_color': '#57A639',
        'section_separator_color': '#6D6552',
        'subtitle_text_color': '#606E8C',
        'text_color': '#6D3F5B',
    },
}



let colorTheme = ''
const setColorTheme = (theme) => {
    colors.telegram = window.Telegram.WebApp.themeParams
    colorTheme = theme
}

const color = (colorName) => {
    console.log(colors, colorTheme, colorName)
    return colors[colorTheme][colorName]
}

export { setColorTheme, color }