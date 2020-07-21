<script lang="typescript" >
    import SimpleBar from '@woden/svelte-simplebar'
    import MainMenu from './common/components/Menu/MainMenu/MainMenu.svelte'
    import routes from "./common/utils/routes";
    import "./localization/localization";

    import { yandexMapsReady, googleMapsReady } from "./stores/maps";

    import Shops from "./pages/Shops/Shops.svelte";
    import Categorization from "./pages/Categorization/Categorization.svelte";

    import { Router, Link, Route } from "svelte-routing";

</script>

<style lang="scss">
    :global(body) {
        margin: 0;
    }

    :global(*) {
        scrollbar-width: none;
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

    :global(.simplebar-scrollbar::before) {
        background: #5b9c35;
    }

    .MainWrapper {
        display: flex;
        flex-direction: column;
        height: 100vh;
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
        <SimpleBar style="max-height: 100vh; width: 100%">
        <div class="RouteWrapper">
                <!-- <Route path="blog/:id" component="{BlogPost}" /> -->
                <Route path={routes.shops} component="{Shops}" />
                <Route path={routes.categorization} component="{Categorization}" />
                <!-- <Route path="/"><Home /></Route> -->
        </div>
        </SimpleBar>
    </div>
</Router>
