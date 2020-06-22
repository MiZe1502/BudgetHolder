<script lang="typescript" >
    import MainMenu from './common/components/MainMenu/MainMenu.svelte'

    import Shops from "./pages/Shops.svelte";

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
