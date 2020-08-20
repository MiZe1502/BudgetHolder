<script lang="ts">


    import {onMount} from "svelte";
    import {setAuthStatus} from "../../stores/auth";
    import {navigate} from "svelte-routing";
    import routes from "../../common/utils/routes"

    onMount(() => {
        setTimeout(() => {
            setAuthStatus(true)
            navigate(routes.budget, { replace: true });
        }, 5000)
    })
</script>

<section>
    AUTH PAGE
</section>