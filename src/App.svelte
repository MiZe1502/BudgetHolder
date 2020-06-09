<script lang="ts" >
    import MainMenu from './common/components/MainMenu/MainMenu.svelte'

    import Shops from "./pages/Shops.svelte";

    import { Router, Link, Route } from "svelte-routing";

</script>

<style lang="scss">

    :global(.Font412Black) {
        font-family: 'Alegreya Sans', sans-serif;
        color: black;
        font-weight: 400;
        font-size: 22px;
    }

    :global(.Font422Black) {
        font-family: 'Alegreya Sans', sans-serif;
        color: black;
        font-weight: 400;
        font-size: 22px;
    }

    :global(body) {
        margin: 0;
    }

    :global(.PrimaryBackground) {
        background-color: #80cf47;
    }

    :global(.PrimaryBackgroundActive) {
        background-color: #5b9c35;
    }

    :global(.FlexHor) {
        display: flex;
        flex-direction: row;
        align-items: center;
    }

    :global(a) {
        text-decoration: none;
        color: black;
    }

    .MainWrapper {
        display: flex;
        flex-direction: column;
    }

</style>

<svelte:head>
  <link href="https://fonts.googleapis.com/css?family=Alegreya+Sans" rel="stylesheet">
</svelte:head>

<Router>
    <div class="MainWrapper">
        <MainMenu />
        <!-- <Route path="blog/:id" component="{BlogPost}" /> -->
        <Route path="shops" component="{Shops}" />
        <!-- <Route path="/"><Home /></Route> -->
    </div>
</Router>
