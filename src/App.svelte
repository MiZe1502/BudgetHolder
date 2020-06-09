<script lang="ts" >
    import MainMenu from './common/components/MainMenu/MainMenu.svelte'

    import Shops from "./pages/Shops.svelte";

    import { Router, Link, Route } from "svelte-routing";

</script>

<style lang="scss">

    :global(body) {
        margin: 0;
    }

    .MainWrapper {
        display: flex;
        flex-direction: column;
    }

</style>


<Router>
    <div class="MainWrapper">
        <MainMenu />
        <!-- <Route path="blog/:id" component="{BlogPost}" /> -->
        <Route path="shops" component="{Shops}" />
        <!-- <Route path="/"><Home /></Route> -->
    </div>
</Router>
