<script>
    import windImg from '../assets/images/wind.png'
    import rainImg from '../assets/images/rain.png'
    import wetImg from '../assets/images/waterdrop.png'
    export let weatherInfo
    export let day, month, time
</script>

<div>
    <div class="time-box">
        {day}
        <span class="month">
            {month}
        </span>
        {time}

    </div>

    <div class="temperature-box">
        <!-- <div class="glow-box"></div> -->
        {weatherInfo.temperature} °C
        <!-- <div class="glow-box"></div> -->
    </div>

    <div class="rainly-box">
        <table align="center" class="table-info">
            <tr>
                <td>
                    <div class="img-container">
                        <img alt="wind" src={windImg}/>
                    </div>
                </td>
                <td>
                    <div class="img-container">
                        <img alt="rainly" src={rainImg}/>
                    </div>
                </td>
                <td>
                    <div class="img-container">
                        <img alt="wet" src={wetImg}/>
                    </div>
                </td>
            </tr>
            <tr>
                <td>{weatherInfo.wind} м/с</td>
                <td>{weatherInfo.rainly_chance}%</td>
                <td>{weatherInfo.wet}%</td>
            </tr>
        </table>
        <!-- Влажность воздуха: {weatherInfo.wet}% -->


    </div>


</div>

<style>
    .month {
        font-weight: 600;
        color: rgb(24, 167, 199);
    }

img {
    display: block;
    width: 50px;
    margin: 0 auto;
    object-fit: contain;
    /* -webkit-filter: drop-shadow(0px 0px 7px rgba(69, 181, 255, 0.5));
    filter: drop-shadow(0px 0px 7px  rgba(69, 181, 255, 0.5));  */
}

.img-container {
    padding: 0;
    margin: 10px 20px;
}

.time-box {
    text-align: center;
    font-size: 1.4rem;
    margin: 1rem auto;
    color: whitesmoke;
    /* -webkit-filter: drop-shadow(0px 0px 5px rgba(255, 255, 255, 0.5));
    filter: drop-shadow(0px 0px 5px  rgba(255, 255, 255, 0.5));  */
}

.temperature-box {
    text-align: center;
    font-size: 2.4rem;
    margin: 0 auto;
    color: whitesmoke;
    font-weight: 1000;
    /* -webkit-filter: drop-shadow(0px 0px 7px rgba(255, 255, 255, 0.5));
    filter: drop-shadow(0px 0px 7px  rgba(255, 255, 255, 0.5));  */
}

/* .glow-box {
    border-radius: 50%;
    background-color: #ffffff;
    box-shadow:
        0 0 40px 5px #70b3ff,
        0 0 60px 10px rgb(158, 216, 255),
        0 0 100px 40px rgb(255, 255, 255);
    opacity: 0.1;
    margin: 0 0 0 0;
} */
/* 
@keyframes glow-box-anim {
    0% {
        box-shadow:
        0 0 40px 5px #70b3ff,
        0 0 60px 10px rgb(158, 216, 255),
        0 0 100px 40px rgb(255, 255, 255);
    }
    50% {
        box-shadow:
        0 0 40px 10px #70b3ff,
        0 0 100px 20px rgb(158, 216, 255),
        0 0 150px 70px rgb(255, 255, 255);
    }
    100% {
        box-shadow:
        0 0 40px 5px #70b3ff,
        0 0 60px 10px rgb(158, 216, 255),
        0 0 100px 40px rgb(255, 255, 255);
    }
} */

.rainly-box {
    text-align: center;
    margin: 0 auto;
}

.table-info {
    border-spacing: -5rem 0rem;
    padding-bottom: 2%;
    border-collapse: collapse;
    margin: 10px auto;
    font-size: 1.5rem;
    /* text-shadow: 0px 0px 5px rgba(255, 255, 255, 0.5); */
}

.table-info td {
    padding: 0;
    margin: 0;
}


</style>