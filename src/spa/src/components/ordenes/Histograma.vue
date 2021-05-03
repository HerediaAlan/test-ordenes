<template>
    <div>
        <canvas id="ordenes-chart"></canvas>
    </div>
</template>

<script>
import Chart from "chart.js";
const moment = require("moment");

// https://www.digitalocean.com/community/tutorials/vuejs-vue-chart-js
export default {
    name: "Histograma",
    props: {
        ordenesProp: Array,
    },
    data() {
        return {
            ordenes: this.ordenesProp,
            dias: this.getDays()
        };
    },
    mounted() {
        const ctx = document.getElementById("ordenes-chart");
        new Chart(ctx, {
            type: "bar",
            data: {
                labels: this.getDays(),
                datasets: [
                    {
                        label: "Número de ordenes",
                        data: this.mapOrdenes(),
                        backgroundColor: "rgba(54,73,93,.5)",
                        borderColor: "#36495d",
                        borderWidth: 3,
                    },
                ],
            },
            options: {
                responsive: true,
                lineTension: 1,
            },
        });
    },
    methods: {
        getDays() {
            // https://stackoverflow.com/a/50731233
            const lastThirtyDays = [...new Array(30)].map((i, idx) => moment().startOf("day").subtract(idx, "days"));

            const l = lastThirtyDays.map((date) => {
                return date.format("DD MMM");
            });

            return l;
        },
        mapOrdenes() {
            // https://www.digitalocean.com/community/tutorials/js-finally-understand-reduce
            const lista = this.dias.map((d) => {
                return this.ordenes.reduce((acc, o) => {
                    if (d === moment(o.FechaCreacion).format("DD MMM")) {
                        acc = acc + 1;
                    }

                    return acc;
                }, 0)
            })

            console.log(lista);

            return lista;
        }
    }
};
</script>