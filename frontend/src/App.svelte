<script>
  import { parser } from '../wailsjs/go/models'
  import { DownloadImage, GetWeather, UpdateWeather, GetDay, GetMonth, GetTime } from '../wailsjs/go/main/App.js'
  import Info from './components/Info.svelte'
  import TitleBar from './components/TitleBar.svelte'
  import PreLoader from './components/Preloader.svelte'
  import SyncImg from './assets/images/sync.png'

  let weatherInfo = parser.WeatherInfo.createFrom({})
  let weatherImg = ''
  let infoProps = {}

  async function getStringDate() {
        infoProps.day = await GetDay()
        infoProps.month = await GetMonth()
        infoProps.time = await GetTime()
  }

  async function InitWeather() {
    weatherInfo = await UpdateWeather()
    await DownloadImage('frontend/src/assets/images/weather.jpg')
    weatherImg = '/src/assets/images/weather.jpg'
    await getStringDate()
  }

  async function update_weather() {
    document.getElementById("preloader").classList.replace("preloader--hidden", "preloader")
    await InitWeather()
    document.getElementById("preloader").classList.replace("preloader", "preloader--hidden")
  }

  InitWeather()

</script>

<main>
  <PreLoader/>
  <div class="main-wrapper">
    <TitleBar/>
    <div class="update-box-wrapper">
      <div class="update-box" id="input">
        <img alt="weather img" id="weather-img" src="{weatherImg}">
          <!-- <button class="btn" on:click={update_weather}>Update</button> -->
          <button class="btn" on:click={update_weather}>
            <img id="syncImg" src={SyncImg} alt="sync"/>
          </button>
      </div>
      <svelte:component this={Info} weatherInfo={weatherInfo} day={infoProps.day} month={infoProps.month} time={infoProps.time}/>
    </div>
  </div>
</main>

<style>

  #syncImg {
    width: 30px;
    height: 30px;
  }

  #weather-img {
    display: block;
    width: 75px;
    height: 75px;
    margin: 0 auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
    
    -webkit-filter: drop-shadow(0px 0px 7px rgba(255, 255, 255, 0.25));
    filter: drop-shadow(0px 0px 7px  rgba(255, 255, 255, 0.25)); 
  }

  main {
    /* flex: 1; */
    /* display: flex; */
    justify-content: center;
    align-items: center;
    /* background-color: rgb(255, 0, 0); */
    height: 100%;
  }
/* 
  .helper-wrapper {
    display: relative;
    vertical-align: middle;
    height: 100%;
    padding: 0;
  } */

  .main-wrapper {
    position: relative;
    height: 100%;
    width: 100%;
    text-align: center;
  }

  .update-box-wrapper {
    position: absolute;
    margin: 0 auto;
    top: 50%;
    left: 50%;
    transform: translateY(-50%) translateX(-50%);
  }

  .update-box {
    position: relative;
    width: 100%;
    height: 100%;
    padding: 0 0;
    margin: auto;
  }

  .update-box .btn {
    top: 20%;
    left: 60%;
    width: 35px;
    height: 35px;
    position: absolute;
    background-color: transparent;
    border: none;
    margin: 0 0;
    padding: 0 0px;
    cursor: pointer;
    rotate: 0deg;

    transition: all 1s ease-in-out;
    --webkit-transition: all 1s ease-in-out;

    animation: rotating 4s linear infinite;
    -webkit-animation: rotating 4s linear infinite;
    animation-play-state: paused;

  }

  .update-box .btn:hover {
    animation-play-state: running;
  }

  .update-box .btn:active {
    rotate: -360deg;
  }

  @keyframes rotating {
    from {
    -webkit-transform: rotate(0deg);
  }
  to {
    -webkit-transform: rotate(-360deg);
  }
  }

</style>
