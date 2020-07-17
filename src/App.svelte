<script lang="typescript" >
    import MainMenu from './common/components/Menu/MainMenu/MainMenu.svelte'

    import { yandexMapsReady, googleMapsReady } from "./stores/maps";

    import Shops from "./pages/Shops/Shops.svelte";

    import { Router, Link, Route } from "svelte-routing";

</script>

<style lang="scss">
    :global(body) {
        margin: 0;
    }

    :global(.PrimaryBackground) {
        background-color: #80cf47;
    }

    :global(a) {
        text-decoration: none;
        color: black;
    }

    :global(button::-moz-focus-inner) {
        border: 0;
    }

    .MainWrapper {
        display: flex;
        flex-direction: column;
    }

    .RouteWrapper {
        padding-top: 100px;
    }

</style>

<svelte:head>
    <script src='https://api-maps.yandex.ru/2.1/?lang=ru_RU&amp;apikey=cf1b8beb-bb0c-4563-9d28-c603002dd2ad'
            type="text/javascript" on:load={() => yandexMapsReady.set(true)}>
    </script>
    <script src="https://maps.googleapis.com/maps/api/js?key=AIzaSyBh9VSJdpY9bDn3OuTaL8i2jvKKKDdPLxc" on:load={() => googleMapsReady.set(true)}>
	</script>
    <link href="https://fonts.googleapis.com/css2?family=Noto+Sans+JP" rel="stylesheet">
</svelte:head>

<Router>
    <div class="MainWrapper">
        <MainMenu />
        <div class="RouteWrapper">
            <!-- <Route path="blog/:id" component="{BlogPost}" /> -->
            <Route path="shops" component="{Shops}" />
            <!-- <Route path="/"><Home /></Route> -->
        </div>
    </div>
</Router>
