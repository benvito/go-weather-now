<script>
    // document.getElementById("preloader").classList.add("preloader--hidden")
</script>

<div id="preloader" class="preloader--hidden">
    <div class="loadingio-spinner-ripple-nq4q5u6dq7r"><div class="ldio-x2uulkbinbj">
        <div></div><div></div>
        </div></div>
</div>

<style>
    @keyframes preload-anim {
        0% {
            top: 102.72px;
            left: 102.72px;
            width: 0;
            height: 0;
            opacity: 1;
        }
        100% {
            top: 6.42px;
            left: 6.42px;
            width: 192.60000000000002px;
            height: 192.60000000000002px;
            opacity: 0;
        }
    }
    
    .ldio-x2uulkbinbj div {
        position: absolute;
        border-width: 4.28px;
        border-style: solid;
        opacity: 1;
        border-radius: 50%;
        animation: preload-anim 1.3s cubic-bezier(0,0.2,0.8,1) infinite;
    }

    .ldio-x2uulkbinbj div:nth-child(1) {
        border-color: #ffffff;
        animation-delay: 0s;
    }

    .ldio-x2uulkbinbj div:nth-child(2) {
        border-color: #46dff0;
        animation-delay: -0.7575757575757576s;
    }

    .loadingio-spinner-ripple-nq4q5u6dq7r {
        width: 214px;
        height: 214px;
        display: inline-block;
        overflow: hidden;
        background: none;
    }

    .ldio-x2uulkbinbj {
        width: 100%;
        height: 100%;
        position: relative;
        transform: translateZ(0) scale(1);
        backface-visibility: hidden;
        transform-origin: 0 0;
    }

    .ldio-x2uulkbinbj div { 
        box-sizing: content-box; 
    }

    #preloader {
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, 0.7);
        position: absolute;
        align-content: center;
        z-index: 1000;
        backdrop-filter: blur(5px)
    }

    .preloader--hidden {
        display: none;
        z-index: -1;
    }
</style>