<script lang="ts">
    import {onMount} from "svelte";
    import {setAuthStatus} from "./auth";
    import {navigate} from "svelte-routing";
    import routes from "../../common/utils/routes"
    import AuthForm from "./AuthForm/AuthForm.svelte";

    onMount(() => {
        setTimeout(() => {
            // setAuthStatus(true)
           // navigate(routes.budget, {replace: true});
        }, 5000)
    })
</script>

<section>
    <AuthForm />
</section>