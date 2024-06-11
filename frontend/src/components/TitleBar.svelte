<script>
    import Logo from '../assets/images/logo.png'
    import * as runtime from '../../wailsjs/runtime/runtime.js'
    import exitImg from '../assets/images/exit.png'
    import maximizeImg from '../assets/images/maximize.png'
    import minimizeImg from '../assets/images/minimize.png'

    function hideApp() {
        runtime.WindowMinimise()
    }

    async function maximizeApp() {
        let maximized = await runtime.WindowIsMaximised()
        if (maximized === true) {
            await runtime.WindowUnmaximise()
        } else {
            await runtime.WindowMaximise()
        }
    }

    function closeApp() {
        runtime.Quit()
    }
</script>

<div class="title-wrapper">
    <div style="--wails-draggable:drag" class="title-bar">
        <ul class="title-bar actions">
            <li>
                <div class="title-bar logo"> 
                    <img id="logo" src="{Logo}" alt="logo">
                    Go Weather Now
                </div>
            </li>
            <li>
                <div class="title-bar buttons">
                    <ul class="title-bar buttons">
                        <li>
                            <button class="action-btn" on:click={hideApp}>
                                <img src="{minimizeImg}" alt="minimize"/>
                            </button>
                        </li>
                        
                        <li>
                            <button class="action-btn" on:click={closeApp}>
                                <img src="{exitImg}" alt="exit"/>
                            </button>
                        </li>
            
                    </ul>
                </div>

            </li>

        </ul>

    </div>

    <!-- <li>
                            <button class="action-btn" on:click={maximizeApp}>
                                <img src="{maximizeImg}" alt="maximize"/>
                            </button>
                        </li> -->
</div>

<style>
    .title-wrapper {
        background-color: rgba(22, 22, 22, 0.575);
    }
    .title-bar {
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 100%;
        height: 30px;
        padding: 0 0;
        z-index: 2;
    }

    ul.title-bar.actions {
        /* background: white; */
        display: flex;
        margin: 0 10px;
        padding: 0;
    }

    ul {
        list-style-type: none;
    }

    .title-bar.logo {
        font-weight: 100;
        font-size: 15px;
        text-align: center;
        color: rgba(255, 255, 255, 0.815);
    }

    #logo {
        width: 20px;
        height: auto;
        margin-right: 7px;
        opacity: 0.8;
    }

    .title-bar.buttons {
        display: flex;
        align-items: center;
    }

    ul.title-bar.buttons {
        display: flex;
        margin-left: auto;
        padding: 0px;
    }

    .title-bar.buttons li {
        /* background: wheat; */

        display: flex;
        margin: 0 auto;
        padding: 0px 0px;
        vertical-align: middle;
        margin-left: 15px;        
    }

    button {
        height: 30px;
        width: 30px;
        display: inline-flex;
        background-color: transparent;
        border: none;
        align-items: center;
        cursor: pointer;
        padding: 0 10px;
        margin: 0;
    }

    button:hover {
        background-color: rgba(255, 255, 255, 0.164);
    }

    .action-btn img {
        filter: invert(1);
        width: 10px;
        height: 10px;
    }
</style>