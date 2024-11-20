tsParticles.load("particles-js", {
    particles: {
        number: {
            value: 80,
            density: {
                enable: true,
                value_area: 800
            }
        },
        shape: {
            type: "circle",
            stroke: {
                width: 0,
                color: "#fff"
            },
        },
        size: {
            value: 4,
            random: true,
            animation: {
                speed: 4,
                minimumValue: 0.1
            }
        },
        opacity: {
            value: 0.5,
            random: true,
            animation: {
                speed: 1,
                minimumValue: 0.1
            }
        },
        links: {
            enable: true,
            distance: 150,
            color: "#fff",
            opacity: 0.4,
            width: 1
        },
        move: {
            enable: true,
            speed: 2,
            direction: "none",
            random: false,
            straight: false,
            outModes: {
                default: "out"
            }
        },
    },
    detectRetina: true,
});