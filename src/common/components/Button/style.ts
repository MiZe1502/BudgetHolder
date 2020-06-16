import { css } from 'emotion';
import colors from "../../styles/colors";

export { Font316Black } from "../../styles/fonts";

export const Button = css`
	width: auto;
	height: auto;
	background-color: ${colors.greenPrimary}; 
	text-decoration: none;
	border: none;
	outline: none;
	padding: 8px;
	border-radius: 4px;
	transition: background-color 300ms ease;
	box-shadow: 0 3px 1px -1px rgba(0, 84, 72, 0.3);

	&:hover,
	&:active {
		background-color: ${colors.greenActive}; 
        cursor: pointer;
        box-shadow: 0 3px 1px -1px rgba(0, 84, 72, 0.15);
	}

	&:active {
		background-color: ${colors.greenLight};
        box-shadow: none;
	}

	&:disabled {
		background-color: ${colors.darkGray};
        border: none;
        cursor: default;
        box-shadow: none;
        color: whitesmoke;
	}
`;